import { useRouter } from "next/dist/client/router";
import Head from "next/head";
import {
  MeDocument,
  useRegisterMutation,
  UserInput,
} from "../generated/graphql";
import { useForm } from "../hooks/useForm";
import Input from "../components/Input";

export default function Register() {
  const initialState = {
      name: "",
    email: "",
    password: "",
  };

  const { onChange, onSubmit, values } = useForm<UserInput>(
    handleRegister,
    initialState
  );

  const [register, { data, loading, error }] = useRegisterMutation({
    update(cache, { data }) {
      cache.writeQuery({
        query: MeDocument,
        data: {
          me: data?.register,
        },
      });
    },
  });

  const router = useRouter();

  async function handleRegister() {
    try {
      const response = await register({ variables: { data: values } });
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
        <title>Grindlists - Register</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <h1 className="text-2xl font-sans">Register</h1>

      <form onSubmit={onSubmit} className="w-96">
          <div className="relative mb-5">
              <Input label="Name" name="name" onChange={onChange}></Input>
          </div>
        <div className="relative mb-5">
          <Input label="E-Mail" name="email" onChange={onChange}></Input>
        </div>
        <div className="relative mb-5">
          <Input
            label="Password"
            name="password"
            onChange={onChange}
            password
          ></Input>
        </div>
        <button
          type="submit"
          className="p-2 border border-steel-700 bg-steel-500 text-gray-100 w-full rounded-md"
        >
          Register
        </button>
      </form>
    </div>
  );
}
