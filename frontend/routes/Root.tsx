import React from "react";
import { Link } from "wouter";

export const Root: React.FC<unknown> = () => {
    return (
        <>
            <h1>Mousetoria</h1>
            <nav>
                <Link href="/app">Continue</Link>
                <Link href="/app">New Game</Link>
                <Link href="/settings">Settings</Link>
                <Link href="/credits">Credits</Link>
                <Link href="/exit">Exit</Link>
            </nav>
        </>
    );
};
