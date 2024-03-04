import React, { useEffect, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import { ToastContainer } from "react-toastify";

import SearchBar from "../components/SearchBar/SearchBar";
import File from "../components/search/File";
import Dataset from "../components/search/Dataset";
import doAdvancedSearch from "../helpers/api/search/advancedSearch";

const AdvancedSearchPage = () => {
	const [datasetResults, setDatasetResults] = useState([]);
	const [fileResults, setFileResults] = useState([]);

	const location = useLocation();
	const navigate = useNavigate();

	useEffect(() => {
		if (location.state) {
			const options = [];
			location.state.forEach((option) => {
				const newOption = option;
				// eslint-disable-next-line prefer-destructuring
				newOption.field = option.field.split("-/--/-")[1];
				options.push(newOption);
			});
			console.log(options);

			doAdvancedSearch(options).then((response) => {
				console.log(response);
				setDatasetResults(response.datasets || []);
				setFileResults(response.files || []);
			}).catch((error) => {
				console.error(error);
			});
		} else {
			navigate("/");
		}
	}, []);

	return (
		<div className="w-screen h-full bg-offwhite">
			<SearchBar />
			<ToastContainer />
			<div className="flex flex-col items-center">
				<div className="max-w-7xl w-full p-4">
					<h2 className="text-2xl font-bold text-gray-800">Datasets</h2>
					{datasetResults.length === 0
						? <p>No datasets found</p>
						: null}
					{datasetResults.map((dataset) => {
						return (
							<Dataset dataset={dataset} />
						);
					})}
					<h2 className="text-2xl font-bold text-gray-800">Files</h2>
					{fileResults.length === 0
						? <p>No files found</p>
						: null}
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

export default AdvancedSearchPage;
