import React, { useEffect, useState } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";

import SearchBar from "../components/SearchBar/SearchBar";
import doSimpleSearch from "../helpers/api/search/simpleSearch";
import File from "../components/search/File";
import Dataset from "../components/search/Dataset";

const SimpleSearchPage = () => {
	const [datasetResults, setDatasetResults] = useState([]);
	const [fileResults, setFileResults] = useState([]);

	const [searchParams] = useSearchParams();

	const navigate = useNavigate();

	useEffect(() => {
		if (searchParams.get("query") === "") {
			navigate("/");
		}
		doSimpleSearch(searchParams.get("query")).then((results) => {
			setDatasetResults(results.datasets);
			setFileResults(results.files);
		});
	}, [searchParams]);

	return (
		<div className="w-screen h-full bg-offwhite">
			<SearchBar />
			<div className="flex flex-col items-center">
				<div className="max-w-7xl w-full p-4">
					<h2 className="text-2xl font-bold text-gray-800">Datasets</h2>
					{datasetResults.map((dataset) => {
						return (
							<Dataset dataset={dataset} />
						);
					})}
					<h2 className="text-2xl font-bold text-gray-800">Files</h2>
					{fileResults.map((file) => {
						return (
							<File file={file} />
						);
					})}
				</div>
			</div>
		</div>
	);
};

export default SimpleSearchPage;
