import React from "react";
import ReactDOM from "react-dom/client";
import { RouterProvider, createBrowserRouter } from "react-router-dom";

import "./index.css";
import "react-toastify/dist/ReactToastify.css";

import reportWebVitals from "./reportWebVitals";
import IndexPage from "./pages/IndexPage/IndexPage";
import LoginPage from "./pages/LoginPage";
import AuthPage from "./pages/AuthPage";
import DatasetPage from "./pages/DatasetPage";
import FilePage from "./pages/FilePage";
import SimpleSearchPage from "./pages/SimpleSearchPage";
import ErrorPage from "./pages/ErrorPage";
import MyDataPage from "./pages/MyDataPage";
import AdvancedSearchPage from "./pages/AdvancedSearchPage";

const router = createBrowserRouter([
	{
		path: "/",
		element: <IndexPage />,
	},
	{
		path: "/login",
		element: <LoginPage />,
	},
	{
		path: "/auth",
		element: <AuthPage />,
	},
	{
		path: "/dataset/:datasetId",
		element: <DatasetPage />,
	},
	{
		path: "/file/:fileId",
		element: <FilePage />,
	},
	{
		path: "/datasets",
		element: <MyDataPage />,
	},
	{
		path: "/search",
		element: <SimpleSearchPage />,
	},
	{
		path: "advancedSearch",
		element: <AdvancedSearchPage />,
	},
	{
		path: "*",
		element: <ErrorPage />,
	},
]);

ReactDOM.createRoot(document.getElementById("root")).render(
	<RouterProvider router={router} />,
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
