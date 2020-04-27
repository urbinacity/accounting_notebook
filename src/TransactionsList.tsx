import React from 'react';
import { TransactionsListState, LoadTransactions } from './TransactionActions';

class TransactionsList extends React.Component<{}, TransactionsListState> {
    constructor(props: {}) {
        super(props);
        this.state = {
            transactions: [],
            activeRows: [],
        }
    }
    componentDidMount() {
        this.loadTransactions();
    }

    async loadTransactions() {
        const response: any = await LoadTransactions();
        this.setState({
            transactions: response,
        });
    }

    formatDate(dateString: Date) {
        const date = new Date(dateString);

        return `Date: ${date.toLocaleDateString()} ${date.toLocaleTimeString()}`
    }

    toggleDetailsRow(transactionId: number) {
        let activeRows = this.state.activeRows;

        if(activeRows.includes(transactionId)) {
            activeRows = activeRows.filter((rowId) => rowId != transactionId)
        }
        else {
            activeRows.push(transactionId);
        }

        return this.setState({
            activeRows,
        });

    }

    render() {
        const { transactions, activeRows } = this.state;
        return (
            <div className="jumbotron">
                <h3>Welcome back!</h3>
                <p className="lead">Here you can see your transactions.</p>
                <hr className="my-4" />
                <br/>
                <table className="table table-hover transactions-list">
                    <thead>
                        <tr className="table-primary">
                            <th scope="col">Type</th>
                            <th scope="col">Amount</th>
                        </tr>
                    </thead>
                    <tbody>
                        {
                            transactions.map((transaction) => (
                                <React.Fragment key={transaction.id}>
                                    <tr className="primary-row" onClick={() => this.toggleDetailsRow(transaction.id)}>
                                        <th scope="row">{transaction.type}</th>
                                        <td className={
                                            transaction.type === 'credit' ? 'text-success' : 'text-danger'
                                        }>{transaction.amount}</td>
                                    </tr>
                                    <tr className={`table-dark details-row ${activeRows.includes(transaction.id) ? 'active' : ''}`}>
                                        <td>
                                            ID:&nbsp;
                                            {transaction.id}
                                        </td>
                                        <td>
                                            {
                                                this.formatDate(transaction.effectiveDate)
                                            }
                                        </td>
                                    </tr>
                                </React.Fragment>
                            ))
                        }
                    </tbody>
                </table>
            </div>
        );
    }
}

export default TransactionsList;
