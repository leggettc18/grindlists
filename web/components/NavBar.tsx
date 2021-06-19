import React from "react";
import Link from "next/link";
import {
  MeDocument,
  useMeQuery,
  useRefreshMutation,
} from "../generated/graphql";
import { NetworkStatus } from "@apollo/client";

export default function NavBar() {
  const { data, loading, error, refetch, networkStatus } = useMeQuery({
    pollInterval: 10000,
    notifyOnNetworkStatusChange: true,
  });
  const [refresh, { data: rData, loading: rLoading, error: rError }] =
    useRefreshMutation({
      notifyOnNetworkStatusChange: true,
      update(cache, { data: refresh }) {
        cache.writeQuery({
          query: MeDocument,
          data: {
            me: refresh ? refresh.refresh : null,
          },
        });
      },
    });

  let body = null;
  const authBody = (
    <>
      <div className="text-gray-100">{data?.me.name}</div>
      <Link href="/logout">
        <a className="hover:bg-sunset-600 text-sunset-100 bg-sunset-500 border border-sunset-600 rounded-lg p-1 shadow-xl">
          Logout
        </a>
      </Link>
    </>
  );
  const noAuthBody = (
    <>
      <Link href="/login">
        <a className="bg-olive-300 text-olive-700 p-1 shadow-xl border-olive-400 border rounded-lg">
          Login
        </a>
      </Link>
      <Link href="/register">
        <a className="bg-seagreen-300 text-seagreen-700 p-1 shadow-xl border-seagreen-400 border rounded-lg">
          Register
        </a>
      </Link>
    </>
  );

  const navbar = (body: JSX.Element | null) => {
    return (
      <div className="flex-shrink">
        <nav className="bg-steel-500">
          <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <div className="flex items-center justify-between h-12">
              <div className="flex items-center">
                <div className="flex-shrink-0">Placeholder</div>
                <div className="hidden md:block">
                  <div className="ml-10 flex items-baseline space-x-4">
                    <Link href="/">
                      <a className="hover:bg-steel-700 text-gray-100 px-3 py-2 rounded-md text-sm font-medium">
                        Dashboard
                      </a>
                    </Link>
                    <Link href="/lists">
                      <a className="text-steel-200 hover:bg-steel-700 hover:text-steel-300 px-3 py-2 rounded-md text-sm font-medium">
                        Lists
                      </a>
                    </Link>
                  </div>
                </div>
              </div>
              <div className="flex space-x-2 items-center">{body}</div>
            </div>
          </div>
        </nav>
      </div>
    );
  };

  if (loading && networkStatus !== NetworkStatus.poll) {
    return navbar(null);
  }
  if (error) {
    refresh()
      .then(() => {
        refetch();
      })
      .catch(() => {
        return navbar(noAuthBody);
      });
  }
  if (!error && data) {
    return navbar(authBody);
  }
  return navbar(noAuthBody);
}
