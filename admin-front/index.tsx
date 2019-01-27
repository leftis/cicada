import * as React from "react";
import * as ReactDOM from "react-dom";

import './styles/admin.scss';

import { Admin } from "./containers/Admin";

ReactDOM.render(
    <Admin compiler="TypeScript" framework="React" />,
    document.getElementById("admin-entry")
);
