import { useQuery, gql } from "@apollo/client";
import styles from "../styles/Home.module.css";
import {HeartIcon} from '@heroicons/react/solid';

const QUERY = gql`
    query Lists {
        lists {
            id
            name
            hearts {
                count
            }
            user {
                name
            }
        }
    }
`;

interface List {
    id: number,
    name: string,
    user: User,
    hearts: Heart,
}

interface Heart {
    count: number
}

interface User {
    name: string,
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
            {data.lists.map((list: List, index: number) => {
                if (index < 4) {
                return <div className={styles.card} key={list.id}>
                    <h3 className="text-blue-500">{list.name}</h3>
                    <p>{list.user.name} {list.hearts.count} <HeartIcon className="h-7 w-7 inline text-blush-500" /></p>
                </div>
                }
            })}
        </div>
    );
}