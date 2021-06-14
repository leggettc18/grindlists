import React, { useState } from "react";
import { Transition } from "@headlessui/react";
import Link from "next/link";

export default function NavBar() {
  return (
    <div>
      <nav className="bg-blue-500">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex items-center justify-between h-12">
            <div className="flex items-center">
              <div className="flex-shrink-0">Placeholder</div>
              <div className="hidden md:block">
                <div className="ml-10 flex items-baseline space-x-4">
                  <Link href="/">
                    <a
                      href="#"
                      className="hover:bg-blue-700 text-gray-100 px-3 py-2 rounded-md text-sm font-medium"
                    >
                      Dashboard
                    </a>
                  </Link>
                  <Link href="/lists">
                    <a className="text-blue-200 hover:bg-blue-700 hover:text-blue-300 px-3 py-2 rounded-md text-sm font-medium">
                      Links
                    </a>
                  </Link>
                </div>
              </div>
            </div>
            <div className="flex space-x-2">
              <Link href="/login">
                <a className="bg-yellow-300 text-yellow-700 p-1 shadow-xl border-yellow-400 border rounded-lg">
                  Login
                </a>
              </Link>
              <Link href="/register">
                <a className="bg-yellow-500 text-yellow-800 p-1 shadow-xl border-yellow-600 border rounded-lg">
                  Register
                </a>
              </Link>
            </div>
          </div>
        </div>
      </nav>
    </div>
  );
}
