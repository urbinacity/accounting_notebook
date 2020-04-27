import React from 'react';
import axios from 'axios';

export type Transaction = {
    id: number,
    type: string,
    amount: number,
    effectiveDate: Date,
}

export type TransactionsListState = {
    transactions: (Transaction)[],
    activeRows: (number)[],
}

export async function LoadTransactions() {
    const url: string = '/api/transactions';

    const response = await axios.get(url)
        .then(function(response: any) {
            return response.data;
        })
        .catch(function(error: any) {
            console.log(error);

            return [];
        });

    return response;
}