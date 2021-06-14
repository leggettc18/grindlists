import { ApolloClient, InMemoryCache } from "@apollo/client";

const client = new ApolloClient({
  uri: "http://api.local/graphql",
  cache: new InMemoryCache(),
});

export default client;
