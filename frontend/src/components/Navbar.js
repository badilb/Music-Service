import React from 'react'
import { BrowserRouter as Router, Switch, Route, NavLink } from 'react-router-dom';

import 'bootstrap/dist/css/bootstrap.min.css'
class Navbar extends React.Component{

    render() {
        return(
            (
                <nav className="navbar navbar-expand-lg navbar-dark bg-dark">
                    <div className="container">
                        <NavLink className="navbar-brand" to="/"> Gingha </NavLink>
                        <button className="navbar-toggler" type="button" data-toggle="collapse"
                                data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false"
                                aria-label="Toggle navigation">
                            <span className="navbar-toggler-icon"></span>
                        </button>
                        <div className="collapse navbar-collapse" id="navbarNav">
                            <ul className="navbar-nav ml-auto">
                                <li className="nav-item active">
                                    <NavLink to={"/"} className={"nav-link"}>Home</NavLink>
                                </li>
                                <li className="nav-item">
                                    <NavLink to={"/profile"} className={"nav-link"}>Profile</NavLink>
                                </li>
                                <li className="nav-item">
                                    <NavLink to={"/login"} className={"nav-link"}>Login</NavLink>
                                </li>
                                <li className="nav-item">
                                    <NavLink to={"/creator"} className={"nav-link"}>Creator</NavLink>
                                </li>

                                <form className="form-inline my-2 my-lg-0">
                                    <input className="form-control mr-sm-2" type="search" placeholder="Search"
                                           aria-label="Search"/>
                                </form>
                            </ul>
                        </div>
                    </div>
                </nav>
            )
        )
    }
}

export default Navbar;
