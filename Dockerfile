FROM node:alpine AS build
WORKDIR /app

COPY package.json package-lock.json ./

RUN npm install

COPY . ./

# check code style first
RUN npm run-script format:check && npm run-script format:write && npm run-script lint:check && npm run-script lint:fix

# run dev
CMD ["npm", "run", "start"]
