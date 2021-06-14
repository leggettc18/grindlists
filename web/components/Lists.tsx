import { useQuery, gql } from "@apollo/client";
import styles from "../styles/Home.module.css";

const QUERY = gql`
    query Lists {
        lists {
            name
            user {
                name
            }
        }
    }
`;

interface List {
    name: String,
    user: User,
}

interface User {
    name: String,
}

export default function Lists() {
    const { data, loading, error } = useQuery(QUERY);

    if (loading) {
        return <h2>loading...</h2>;
    }

    if (error) {
        console.error(error);
        return null;
    }

    return (
        <div className={styles.grid}>
            {data.lists.map((list: List) => (
                <div className={styles.card}>
                    <h3>{list.name}</h3>
                    <p>{list.user.name}</p>
                </div>
            ))}
        </div>
    );
}