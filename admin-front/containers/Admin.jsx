import * as React from "react";

import { withStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import IconButton from '@material-ui/core/IconButton';

// 'HelloProps' describes the shape of props.
// State is never set so we use the '{}' type.
export class Admin extends React.Component() {
    render() {
        return (
            <AppBar position="static">
                <Toolbar>
                    <IconButton color="inherit" aria-label="Menu">
                    </IconButton>
                    <Typography variant="h6" color="inherit" >
                        News
                    </Typography>
                    <Button color="inherit">Login</Button>
                </Toolbar>
            </AppBar>
        );
    }
}