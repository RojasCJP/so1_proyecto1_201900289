import React from 'react'
import { Link } from 'react-router-dom'

export class Navbar extends React.Component {
    render() {
        return (
            <div>
                <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
                    <div class='container'>
                        <a class="navbar-brand" href='/'>Proyecto 1</a>
                        <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
                            <span class="navbar-toggler-icon"></span>
                        </button>
                        <div class="collapse navbar-collapse" id="navbarSupportedContent">
                            <ul class="navbar-nav mr-auto">
                                <li class="nav-item active">
                                        <a class="nav-link" href='/'>Datos</a>
                                </li>
                                <li class="nav-item active">
                                        <a class="nav-link" href='cpu'>CPU</a>
                                </li>
                                <li class="nav-item active">
                                        <a class="nav-link" href='memory'>Memoria</a>
                                </li>
                            </ul>
                        </div>
                    </div>
                </nav>
            </div>
        )
    }
}