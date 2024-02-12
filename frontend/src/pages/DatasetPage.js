import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

import getDataset from "../helpers/api/webApi/dataset/getDataset";
import getDatasetFiles from "../helpers/api/webApi/file/getFilesByDataset";
import SearchBar from "../components/SearchBar/SearchBar";
import Loader from "../components/basic/Loader";
import File from "../components/basic/File";
import Attribute from "../components/DatasetPage/Attribute";

const DatasetPage = () => {
	const [dataset, setDataset] = useState({});
	const [files, setFiles] = useState([]);

	const { datasetId } = useParams();

	useEffect(() => {
		getDataset(datasetId).then((data) => {
			setDataset(data);
		});
		getDatasetFiles(datasetId).then((data) => {
			setFiles(data);
		});
	}, []);

	return (
		<div className="w-screen h-full bg-offwhite">
			<SearchBar />
			{ dataset
				? (
					<div className="flex justify-center items-center w-full">
						<div className="w-full p-8 max-w-7xl flex">
							<div className="w-1/2 p-2">
								<h1 className="text-3xl font-bold">{dataset.dataset_name}</h1>
								<div className="h-[2px] w-full bg-oxfordblue mb-4" />
								<Attribute attributeName="Author" value={dataset.owner_name} />
							</div>
							<div className="w-1/2 p-2 flex flex-wrap">
								{files.map((file) => {
									return (
										<File key={file.id} file={file} />
									);
								})}
							</div>
						</div>
					</div>
				)
				: <Loader />}
		</div>
	);
};

export default DatasetPage;
