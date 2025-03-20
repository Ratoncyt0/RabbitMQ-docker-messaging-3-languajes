# Use the official Node.js image
FROM node:latest

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy package.json and package-lock.json (if available) into the container
COPY package*.json ./

# Install the dependencies (including amqplib)
RUN npm install

# Copy the rest of the application code (e.g., dat.js) into the container
COPY dat.js .

# Run the dat.js file with Node.js
CMD ["node", "dat.js"]
