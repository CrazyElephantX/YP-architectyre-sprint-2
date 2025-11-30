#!/bin/bash

# Скрипт для проверки и создания топиков Kafka

echo "Waiting for Kafka to be ready..."
sleep 10

# Проверка существующих топиков
echo "Checking existing topics..."
EXISTING_TOPICS=$(kafka-topics.sh --bootstrap-server kafka:9092 --list)

# Список необходимых топиков
REQUIRED_TOPICS=("movie-events" "user-events" "payment-events")

# Функция для создания топика
create_topic() {
    local topic=$1
    echo "Creating topic: $topic"
    kafka-topics.sh --create --topic $topic --bootstrap-server kafka:9092 --partitions 1 --replication-factor 1
}

# Проверка и создание топиков
for topic in "${REQUIRED_TOPICS[@]}"; do
    if echo "$EXISTING_TOPICS" | grep -q "^${topic}$"; then
        echo "Topic $topic already exists"
    else
        echo "Topic $topic does not exist. Creating..."
        create_topic $topic
    fi
done

echo "All required topics checked/created"