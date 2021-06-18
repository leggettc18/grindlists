import { ApolloError, gql } from "@apollo/client";
import { useRouter } from "next/dist/client/router";
import Head from "next/head";
import {
  LoginInput,
  MeDocument,
  MeQuery,
  useLoginMutation,
  useMeQuery,
} from "../generated/graphql";
import { useForm } from "../hooks/useForm";

export default function Login() {
  const initialState = {
    email: "",
    password: "",
  };

  const { onChange, onSubmit, values } = useForm<LoginInput>(
    handleLogin,
    initialState
  );

  const [login, { data, loading, error }] = useLoginMutation({
    update(cache, { data }) {
      cache.writeQuery({
        query: MeDocument,
        data: {
          me: data?.login,
        },
      });
    },
  });

  const router = useRouter();

  async function handleLogin() {
    try {
      const response = await login({ variables: { data: values } });
      if (response.errors) {
        console.error(error);
      } else {
        router.push("/");
      }
    } catch {
      console.error(error?.graphQLErrors);
    }
  }

  return (
    <div className="flex flex-col items-center space-y-4 pt-4">
      <Head>
        <title>Grindlists - Login</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <h1 className="text-2xl font-sans">Login</h1>

      <form onSubmit={onSubmit}>
        <div className="flex flex-col">
          <label htmlFor="email">E-Mail</label>
          <input
            className="p-2 rounded-lg border-2 border-gray-500 focus:outline-none focus:border-steel-500"
            type="text"
            name="email"
            id="email"
            onChange={onChange}
          />
        </div>
        <div className="flex flex-col">
          <label htmlFor="password">Password</label>
          <input
            className="p-2 rounded-lg border-2 border-gray-500 focus:outline-none focus:border-steel-500"
            type="password"
            name="password"
            id="password"
            onChange={onChange}
          />
        </div>
        <button
          type="submit"
          className="p-2 border border-steel-700 bg-steel-500 text-gray-100"
        >
          Login
        </button>
      </form>
    </div>
  );
}
