import { useQuery, gql } from "@apollo/client";
import styles from "../styles/Home.module.css";

const QUERY = gql`
  query Hello {
    hello
  }
`;

export default function Countries() {
  const { data, loading, error } = useQuery(QUERY);

  if (loading) {
    return <h2>loading...</h2>;
  }

  if (error) {
    console.error(error);
    return null;
  }

  const hellos = [data.hello, data.hello, data.hello, data.hello];

  return (
    <div className={styles.grid}>
      {hellos.map((hello) => (
        <div className={styles.card}>
          <h3>{hello}</h3>
          <p>{hello}</p>
        </div>
      ))}
    </div>
  );
}
