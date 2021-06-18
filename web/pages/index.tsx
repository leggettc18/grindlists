import Head from 'next/head'
import Image from 'next/image'
import styles from '../styles/Home.module.css'

export default function Home() {
  return (
    <div className="h-full flex-grow flex flex-col justify-between content-center">
      <Head>
        <title>Grindlists</title>
        <meta
          name="description"
          content="Helping gamers organize their grinding sessions."
        />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className="h-full flex-grow flex flex-col justify-center items-center">
        <h1 className={styles.title}>
          Welcome to <a href="https://nextjs.org">Grindlists!</a>
        </h1>

        <p className={styles.description}>
          Helping gamers organize their grinding sessions since 2021!
        </p>

        <div className={styles.grid}>
          <a href="https://nextjs.org/docs" className={styles.card}>
            <h2>Documentation &rarr;</h2>
            <p>Find in-depth information about Next.js features and API.</p>
          </a>

          <a href="https://nextjs.org/learn" className={styles.card}>
            <h2>Learn &rarr;</h2>
            <p>Learn about Next.js in an interactive course with quizzes!</p>
          </a>

          <a
            href="https://github.com/vercel/next.js/tree/master/examples"
            className={styles.card}
          >
            <h2>Examples &rarr;</h2>
            <p>Discover and deploy boilerplate example Next.js projects.</p>
          </a>

          <a
            href="https://vercel.com/new?utm_source=create-next-app&utm_medium=default-template&utm_campaign=create-next-app"
            className={styles.card}
          >
            <h2>Deploy &rarr;</h2>
            <p>
              Instantly deploy your Next.js site to a public URL with Vercel.
            </p>
          </a>
        </div>
      </main>
    </div>
  );
}
