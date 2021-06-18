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

      <form onSubmit={onSubmit} className="w-96">
        <div className="relative mb-5">
          <input
            className="peer p-2 border-b border-gray-400 focus:outline-none focus:border-steel-500 w-full p-3 h-16 pt-8"
            type="text"
            name="email"
            id="email"
            placeholder=" "
            onChange={onChange}
          />
          <label
            htmlFor="email"
            className="absolute -top-4 left-0 px-3 py-5 transform origin-left transition-all duration-100 ease-in-out text-gray-400 peer-placeholder-shown:top-0 peer-focus:-top-4 peer-focus:text-steel-500 text-sm peer-focus:text-sm peer-placeholder-shown:text-lg h-full"
          >
            E-Mail
          </label>
        </div>
        <div className="relative mb-5">
          <input
            className="peer p-2 border-b border-gray-400 focus:outline-none focus:border-steel-500 w-full p-3 h-16 pt-8"
            type="password"
            name="password"
            id="password"
            placeholder=" "
            onChange={onChange}
          />
          <label
            htmlFor="password"
            className="absolute -top-4 left-0 px-3 py-5 transform origin-left transition-all duration-100 ease-in-out text-gray-400 peer-placeholder-shown:top-0 peer-focus:-top-4 peer-focus:text-steel-500 text-sm peer-focus:text-sm peer-placeholder-shown:text-lg h-full"
          >
            Password
          </label>
        </div>
        <button
          type="submit"
          className="p-2 border border-steel-700 bg-steel-500 text-gray-100 w-full rounded-md"
        >
          Login
        </button>
      </form>
    </div>
  );
}
