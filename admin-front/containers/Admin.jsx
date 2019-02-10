import React from 'react'
import {
    BrowserRouter as Router,
    Route,
} from 'react-router-dom'

import AdminLoginPage from "../components/AdminLoginPage";

export class Admin extends React.Component {
    render() {
        return (
            <Router>
                <div>
                    <Route path="/admin/login" component={AdminLoginPage}/>
                </div>
            </Router>
        );
    }
}