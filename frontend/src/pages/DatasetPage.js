import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { ToastContainer } from "react-toastify";

import getDataset from "../helpers/api/webApi/dataset/getDataset";
import getDatasetFiles from "../helpers/api/webApi/file/getFilesByDataset";
import SearchBar from "../components/SearchBar/SearchBar";
import Loader from "../components/basic/Loader";
import File from "../components/basic/File";
import getDatasetAttributes from "../helpers/api/webApi/datasetAttributes/getDatasetAttributes";
import Attributes from "../components/DatasetPage/Attribute";
import Upload from "../components/IndexPage/Upload";
import ErrorToast from "../helpers/toast/errorToast";

const DatasetPage = () => {
	const [dataset, setDataset] = useState({});
	const [files, setFiles] = useState([]);
	const [attributes, setAttributes] = useState([]);

	const [refreshAttribute, setRefreshAttribute] = useState(false);

	const { datasetId } = useParams();

	useEffect(() => {
		if (refreshAttribute) {
			getDatasetAttributes(datasetId).then((data) => {
				setAttributes(data);
			}).catch((error) => {
				ErrorToast(error);
			});
			setRefreshAttribute(false);
		}
	}, [refreshAttribute]);

	useEffect(() => {
		getDataset(datasetId)
			.then((data) => {
				setDataset(data);
			}).catch((error) => {
				ErrorToast(error);
			});
		getDatasetFiles(datasetId)
			.then((data) => {
				setFiles(data);
			}).catch((error) => {
				ErrorToast(error);
			});
		getDatasetAttributes(datasetId)
			.then((data) => {
				setAttributes(data);
			}).catch((error) => {
				ErrorToast(error);
			});
		setRefreshAttribute(false);
	}, []);

	return (
		<div className="w-screen h-full bg-offwhite">
			<SearchBar />
			<ToastContainer />
			{ dataset
				? (
					<div className="flex justify-center items-center w-full">
						<div className="w-full p-8 max-w-7xl flex">
							<div className="w-1/2 p-2">
								<h1 className="text-3xl font-bold">{dataset.dataset_name}</h1>
								<div className="h-[2px] w-full bg-oxfordblue mb-4" />
								<div className="w-full flex">
									<p className="text-gray-800 mr-4">Author: </p>
									<p className="font-medium">{dataset.owner_name}</p>
								</div>

								<h2 className="mt-6 text-2xl font-semibold">Dataset Attributes</h2>
								<div className="w-full max-h-64">
									<Attributes
										attributes={attributes}
										datasetId={datasetId}
										setNeedRefresh={setRefreshAttribute}
									/>
								</div>
							</div>
							<div className="w-1/2 p-2">
								<div className="w-full flex flex-wrap">
									{files.map((file) => {
										return (
											<File key={file.id} file={file} />
										);
									})}
								</div>
								<div className="w-full h-64 mt-8">
									<Upload
										datasetId={datasetId}
										onFinish={() => {
											getDatasetFiles(datasetId)
												.then((data) => {
													setFiles(data);
												}).catch((error) => {
												ErrorToast(error);
											});
										}}
									/>
								</div>
							</div>
						</div>
					</div>
				)
				: <Loader />}
		</div>
	);
};

export default DatasetPage;
