import React from "react";
import SearchBar from "../../components/SearchBar/SearchBar";
import useAuth from "../../hooks/useAuth";
import LoggedInContent from "./LoggedInContent";

const IndexPage = () => {
	const { username } = useAuth();

	return (
		<div className="w-screen h-full bg-offwhite">
			<SearchBar />
			{username ? <LoggedInContent /> : null}
		</div>
	);
};

export default IndexPage;
