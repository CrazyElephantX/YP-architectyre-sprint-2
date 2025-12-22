FROM node:18-alpine

WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies including newman reporters
RUN npm install

# Copy test files
COPY . .

# Set default command to run tests
ENTRYPOINT ["node", "run-tests.js"]

# Default arguments - can be overridden at runtime
CMD ["--environment", "docker"]