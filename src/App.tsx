import React from 'react';
import Navbar from './Navbar'
import TransactionsList from './TransactionsList'
import 'bootswatch/dist/flatly/bootstrap.min.css';
import './App.css';

function App() {
    return (
        <div>
            <header>
                <Navbar />
            </header>
            <div className="container-sm">
                <TransactionsList />
            </div>
        </div>
    );
}

export default App;
