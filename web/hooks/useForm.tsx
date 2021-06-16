import React, { useState } from "react"

export const useForm = <T extends {}>(callback: any, initialState: T) => {
    const [values, setValues] = useState<T>(initialState);

    const onChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setValues({ ...values, [event.target.name]: event.target.value });
    };

    const onSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        await callback();
    };

    return {
        onSubmit,
        onChange,
        values
    };
}