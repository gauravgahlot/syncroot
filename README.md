# SyncRoot

SyncRoot is a data synchronization service that allows you to manage contacts across multiple CRM providers such as Salesforce and HubSpot. It provides a unified API to create, read, update, and delete contacts, and ensures data consistency across all connected providers.

## Table of Contents

- [Overview](#overview)
- [Architecture](#architecture)
- [Features](#features)
- [Components](#components)
- [Design Decisions](#design-decisions)
- [API Endpoints](#api-endpoints)
- [Request Flow](#request-flow)

## Overview

SyncRoot acts as a middleware between your application and various CRM providers. When you create, update, or delete a contact through SyncRoot's API, the system ensures that the changes are propagated to all connected providers, handling the necessary data transformations and API calls.

## Architecture

The architecture of SyncRoot follows a message-driven approach with separate components for API handling, message processing, and provider integration.

```mermaid
graph TD
    A[API Server] -->|Enqueue| B[Message Queue]
    B --> C[Forwarder Worker]
    C -->|Store| D[(Database)]
    C -->|Transform & Send| E[Provider 1 - Salesforce]
    C -->|Transform & Send| F[Provider 2 - HubSpot]
    C -->|Failed Messages| G[DLQ]
    G --> H[DLQ Handler Worker]

    E -->|Webhooks| J[Webhook Endpoints]
    F -->|Webhooks| J
    J -->|Provider-specific Queue| K[Provider-specific Topics]
    K --> I[Webhook Handler Worker]
    I -->|Store Changes| D
```

- The message queue implementation uses Kafka to provide a robust, scalable, and distributed messaging system. 
- Topic-based queuing allows segregation of messages per provider, enabling efficient processing and easier error isolation.
- The Dead-Letter Queue (DLQ) is used to handle failed messages, ensuring that problematic messages do not block the pipeline and can be retried or examined later.

## Features

- CRUD operations for contacts
- Synchronization with multiple CRM providers (Salesforce, HubSpot)
- Message-based architecture for reliability and scalability
- Dead-letter queue (DLQ) for handling failed messages
- Webhook support for receiving updates from providers
- Configurable via environment variables

## Components

### API Server

Exposes RESTful endpoints for managing contacts and other resources.

### Workers

- **Forwarder Worker**: Processes messages from the queue, stores data in the database, and forwards to providers. Implements retry strategies with exponential backoff to handle transient failures.
- **DLQ Handler Worker**: Processes failed messages from the Dead-Letter Queue with backoff and alerting mechanisms.
- **Webhook Handler Worker**: Processes webhooks from providers that are queued in provider-specific topics. Also supports retry and backoff strategies to ensure reliable processing.

### Database

Stores contact information and synchronization status.

### Providers

Integrations with CRM systems:

- Salesforce
- HubSpot

Adding a new provider requires implementing a new transformer for data format conversion and binding the provider to a new queue/topic for message handling.

### Transformers

Bi-directional transformers convert data between the internal format and provider-specific formats. These transformers are pluggable, allowing easy addition of new providers and custom transformation logic.

### Security

SyncRoot uses bearer token authentication to secure API endpoints. Additionally, webhook endpoints can optionally validate signatures to ensure authenticity and integrity of incoming requests from providers.

### Monitoring & Observability

SyncRoot integrates with Prometheus for metrics collection, Zap for structured logging, and optionally supports OpenTelemetry traces. Alerts can be configured via Grafana or other Prometheus-compatible alert managers.

## Design Decisions

- **Message-driven architecture**: Enables asynchronous processing, scalability, and decoupling between components.
- **Dead-Letter Queue (DLQ) for resilience**: Ensures that failed messages are isolated and can be retried or analyzed without blocking the main processing pipeline.
- **OpenAPI-based validation**: Provides strict input validation to maintain API consistency and prevent malformed requests.
- **Bi-directional transformers**: Facilitate seamless data conversion between internal and provider formats, supporting synchronization in both directions.
- **Topic-per-provider queue design**: Allows fine-grained control over message processing per provider, improving fault isolation and scalability.
- **Kafka for messaging**: Chosen for its durability, scalability, and mature ecosystem supporting high-throughput, fault-tolerant, distributed streaming use cases.

## API Endpoints

| Endpoint               | Method | Description                      |
| ---------------------- | ------ | -------------------------------- |
| `/contacts`            | POST   | Create a new contact             |
| `/contacts/{id}`       | GET    | Retrieve a contact by ID         |
| `/contacts/{id}`       | PUT    | Update an existing contact       |
| `/contacts/{id}`       | DELETE | Delete a contact by ID           |
| `/webhooks/{provider}` | POST   | Receive webhooks from providers  |
| `/webhooks/salesforce` | POST   | Receive webhooks from Salesforce |
| `/webhooks/hubspot`    | POST   | Receive webhooks from HubSpot    |
| `/healthz`             | GET    | Health check endpoint            |

## Request Flow

### Create Contact Flow

1. Client sends a POST request to `/contacts` with contact details
2. API server validates the request using OpenAPI schema
3. Contact is enqueued in the message queue for processing
4. Forwarder worker picks up the message
5. Contact is stored in the database
6. Contact is transformed and sent to each provider (Salesforce, HubSpot)
7. If any provider integration fails, the message is sent to DLQ for retry with backoff
8. The system ensures eventual consistency by retrying failed messages until successful or manual intervention

### Webhook Flow

1. When a CRM provider (Salesforce, HubSpot) has an update, it sends a webhook to the respective endpoint (`/webhooks/salesforce` or `/webhooks/hubspot`)
2. The webhook handler validates the incoming payload and authenticity
3. The webhook data is enqueued in a provider-specific queue topic
4. The Webhook Handler Worker picks up the message from the appropriate queue
5. The worker processes the data, transforming it from provider format to internal format
6. The updated data is stored in the database, ensuring consistency across the system
7. If necessary, the change is propagated to other providers to keep all systems in sync
8. Retry and backoff mechanisms handle transient failures, ensuring eventual consistency
