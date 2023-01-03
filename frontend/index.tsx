import { StoreProvider } from "easy-peasy";
import React from "react";
import { createRoot } from "react-dom/client";
import { Route, Switch } from "wouter";

import { Root } from "./routes/Root";
import { App } from "./App";
import { store } from "./state";

createRoot(document.getElementById("root")!).render(
    <React.StrictMode>
        <StoreProvider store={store}>
            <Switch>
                <Route path="/root" component={Root} />
                <Route path="/app" component={App} />
            </Switch>
        </StoreProvider>
    </React.StrictMode>
);
