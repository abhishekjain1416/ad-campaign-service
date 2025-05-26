# Ad Campaign Delivery Service

A scalable microservice designed to deliver targeted ad campaigns efficiently by integrating a Match Engine and leveraging Redis caching for performance optimization.

---

## Overview

This service fetches targeted campaign IDs from the Match Engine based on different dimensions such as country, OS, and app. It then retrieves detailed campaign data in batches, utilizing Redis caching to minimize database load and reduce latency.

---

## Architecture

### Core Components

- **Match Engine Service**  
  Acts as the targeting logic engine. Given request parameters, it returns a list of campaign IDs that match the targeting criteria.

- **Campaign Service**  
  Responsible for fetching campaign details such as campaign code, image URL, and call-to-action (CTA). Implements Redis cache-aside pattern to improve response times and reduce database hits.

- **Delivery Service**  
  Orchestrates the workflow by:
  1. Querying the Match Engine for targeted campaign IDs.
  2. Batch processing campaign IDs in configurable chunks.
  3. Concurrently fetching campaign details from the Campaign Service.
  4. Aggregating results and returning a consolidated response.

### Redis Caching Strategy

- **Cache-Aside Pattern**  
  On fetching campaign details, the service first queries Redis. If data is missing or stale, it fetches from the database and updates Redis asynchronously.
  
- **Batch Processing & Concurrency**  
  Campaign IDs are processed in configurable batch sizes (default 100) and fetched concurrently using goroutines. This reduces latency and optimizes throughput.

---

## Getting Started

### Prerequisites

- Go 1.20+  
- Redis server  
- PostgreSQL

### Installation

1. Clone the repo:

   ```bash
   gh repo clone abhishekjain1416/ad-campaign-service
   cd ad-campaign-service

2. Setup environment variables (create .env file):

   ```bash
   DATABASE_URL = "host=localhost user=postgres password=postgres dbname=postgres port=5432"
   REDIS_URL = "redis://127.0.0.1:6379"
   REDIS_CLUSTER = "false"
   DELIVERY_CAMPAIGN_BATCH_SIZE = 100

3. Run the service:
   
   ```bash
   go mod vendor
   go run cmd/server/main.go
