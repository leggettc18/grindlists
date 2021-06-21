interface InputProps {
  label: string;
  name: string;
  password?: boolean;
  onChange: (event: React.ChangeEvent<HTMLInputElement>) => void;
}

const Input = (props: InputProps) => {
  return (
    <>
      <input
        className="peer p-2 border-b border-gray-400 focus:outline-none focus:border-steel-500 w-full p-3 h-16 pt-8"
        type={props.password ? "password" : "text"}
        name={props.name}
        id={props.name}
        placeholder=" "
        onChange={props.onChange}
      />
      <label
        htmlFor={props.name}
        className="absolute -top-4 left-0 px-3 py-5 transform origin-left transition-all duration-100 ease-in-out text-gray-400 peer-placeholder-shown:top-0 peer-focus:-top-4 peer-focus:text-steel-500 text-sm peer-focus:text-sm peer-placeholder-shown:text-lg h-full"
      >
        {props.label}
      </label>
    </>
  );
};

export default Input;
