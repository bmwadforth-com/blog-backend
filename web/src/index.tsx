import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import {RecoilRoot} from 'recoil';
import {BrowserRouter, ScrollRestoration} from "react-router-dom";
import { initializeApp } from "firebase/app";
import { getAnalytics, logEvent } from "firebase/analytics";

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
    apiKey: "AIzaSyDU6p63BqTP0FIHqVXpxY4ytrwY6-AXPoE",
    authDomain: "digital-node-1176.firebaseapp.com",
    projectId: "digital-node-1176",
    storageBucket: "digital-node-1176.appspot.com",
    messagingSenderId: "416443939101",
    appId: "1:416443939101:web:e75dc69b02d92040a149f4",
    measurementId: "G-EK6VC80EVT"
};

const app = initializeApp(firebaseConfig);
export const analytics = getAnalytics(app);
logEvent(analytics, 'notification_received');

const root = ReactDOM.createRoot(
    document.getElementById('root') as HTMLElement
);
root.render(
    <React.StrictMode>
        <RecoilRoot>
            <BrowserRouter>
                <ScrollRestoration />
                <App />
            </BrowserRouter>
        </RecoilRoot>
    </React.StrictMode>
);
// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
