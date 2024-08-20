# Step 1: Build the frontend using Vite
FROM node:18-alpine AS build

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

RUN npm run build

# Step 2: Serve the frontend using a lightweight web server
FROM nginx:alpine

# Copy Vite build output to Nginx's default directory
COPY --from=build /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
