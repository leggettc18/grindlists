import Head from "next/head";
import styles from "../styles/Home.module.css";
import ClientOnly from "../components/ClientOnly";
import ListsComponent from "../components/Lists";

export default function Lists() {
    return (
        <div className={styles.container}>
            <Head>
                <title>Grindlists - All Lists</title>
                <link rel="icon" href="/favicon.ico" />
            </Head>

            <main className={styles.main}>
                <h1 className={styles.title}>
                    All Lists
                </h1>
                <ClientOnly>
                    <ListsComponent />
                </ClientOnly>
            </main>
        </div>
    )
}