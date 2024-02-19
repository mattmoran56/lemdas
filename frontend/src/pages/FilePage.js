import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

import SearchBar from "../components/SearchBar/SearchBar";
import getFile from "../helpers/api/webApi/file/getFile";
import Attributes from "../components/FilePage/Attributes";
import getFileAttributes from "../helpers/api/webApi/fileAttributes/getFileAttributes";
import getPreviewURL from "../helpers/api/webApi/file/getPreview";
import Loader from "../components/basic/Loader";

const FilePage = () => {
	const [file, setFile] = useState({});
	const [attributes, setAttributes] = useState([]);
	const [previewUrl, setPreviewUrl] = useState("");

	const [refreshAttribute, setRefreshAttribute] = useState(false);

	const { fileId } = useParams();

	useEffect(() => {
		if (refreshAttribute) {
			getFileAttributes(fileId).then((data) => {
				setAttributes(data);
			});
			setRefreshAttribute(false);
		}
	}, [refreshAttribute]);

	useEffect(() => {
		getFile(fileId).then((data) => {
			setFile(data);
		});
		getFileAttributes(fileId).then((data) => {
			setAttributes(data);
		});
		getPreviewURL(fileId).then((data) => {
			setPreviewUrl(data.url);
		});
	}, []);
	return (
		<div className="w-screen h-full bg-offwhite">
			<SearchBar />
			<div className="flex justify-center items-center w-full">
				<div className="w-full p-8 max-w-7xl w-full flex">
					<div className="w-1/2 p-2">
						<h1 className="text-3xl font-bold">{file.name}</h1>
						<div className="h-[2px] w-full bg-oxfordblue mb-4" />
						<div className="w-full flex">
							<p className="text-gray-800 mr-4">Author: </p>
							<p className="font-medium">{file.owner_name}</p>
						</div>
						<div className="w-full flex">
							<p className="text-gray-800 mr-4">Dataset: </p>
							<p className="font-medium hover:underline">
								<a href={`/dataset/${file.dataset_id}`}>
									{file.dataset_name}
								</a>
							</p>
						</div>

						<h2 className="mt-6 text-2xl font-semibold">File Attributes</h2>
						<div className="w-full max-h-80 overflow-y-auto">
							<Attributes
								attributes={attributes}
								fileId={fileId}
								setNeedRefresh={setRefreshAttribute}
							/>
						</div>
					</div>
					<div className="w-1/2 p-2">
						{file.status !== "processed" && file.status !== "awaitingtxt"
							? <p>Preview will be available when file is finished processing</p>
							: null}
						{(file.status === "processed" || file.status === "awaitingtxt") && previewUrl !== ""
							? (
								<img
									src={previewUrl}
									alt="preview"
								/>
							)
							: <div className="w-full h-full flex justify-center items-center"><Loader /></div> }
					</div>
				</div>
			</div>
		</div>
	);
};

export default FilePage;
