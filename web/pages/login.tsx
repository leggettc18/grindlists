import { useRouter } from "next/dist/client/router";
import Head from "next/head";
import {
  LoginInput,
  MeDocument,
  useLoginMutation,
} from "../generated/graphql";
import { useForm } from "../hooks/useForm";
import Input from "../components/Input";

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
          <Input label="E-Mail" name="email" value={values.email} onChange={onChange}></Input>
        </div>
        <div className="relative mb-5">
          <Input
            label="Password"
            name="password"
            value={values.password}
            onChange={onChange}
            password
          ></Input>
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
