import React from 'react';
import {
    Route,
    Redirect,
    BrowserRouter as Router,
} from 'react-router-dom';

import AdminLoginPage from '../components/AdminLoginPage';
import AdminDashboard from './AdminDashboard';

export class Admin extends React.Component {
    render() {
        return (
            <Router>
                <div>
                    <Route exact path="/admin" render={() => <Redirect to="/admin/dashboard" />}/>
                    <Route path="/admin/dashboard" component={AdminDashboard} />
                    <Route path="/admin/login" component={AdminLoginPage}/>
                </div>
            </Router>
        );
    }
}