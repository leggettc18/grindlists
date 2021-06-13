FROM node:16.3.0-alpine as base
WORKDIR /client
COPY package*.json ./
RUN yarn install --frozen-lockfile && yarn cache clean --force
FROM base as dev
ENV NODE_ENV=development
ENV PATH /client/node_modules/.bin:$PATH
EXPOSE 3000
RUN mkdir /client/app
RUN yarn install --only=development && yarn cache clean --force
RUN yarn config list
WORKDIR /client/app
CMD ["yarn", "dev"]
FROM dev as test
COPY . .
RUN yarn audit
RUN CI=true yarn test --env=jsdom
FROM test as build-stage
RUN yarn build
FROM aquasec/trivy:0.18.3 as trivy
RUN trivy nginx:1.17-alpine && \
    echo "No image vulnerabilities" > result
FROM nginx1.17-alpine as prod
COPY --from=trivy result secure
COPY --from=build-stage /client/app/build /usr/share/nginx/html
COPY --from=build-stage /client/app/nginx /etc/nginx/
EXPOSE 80
HEALTHCHECK CMD ["wget", "-q", "0.0.0.0:80"]