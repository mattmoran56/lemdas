import React, { useEffect, useState } from "react";
import { ArrowRightIcon } from "@heroicons/react/24/solid";

import Button from "../../basic/Button";
import uploadFile from "../../../helpers/api/upload/uploadFile";
import UploadFile from "./UploadFile";
import ChooseDataset from "./ChooseDataset";
import getDatasets from "../../../helpers/api/webApi/dataset/getDatasets";
import LoadingUpload from "./LoadingUpload";

const Upload = ({ datasetId, onFinish }) => {
	const [file, setFile] = useState(null);
	const [dataset, setDataset] = useState(datasetId);
	const [isPublic, setIsPublic] = useState(false);

	const [uploaded, setUploaded] = useState(false);
	const [loading, setLoading] = useState(false);

	const [datasets, setDatasets] = useState([]);

	const handleFinish = () => {
		setLoading(true);
		uploadFile(file, dataset, isPublic).then(() => {
			setFile(null);
			setDataset(null);
			setUploaded(true);
			setLoading(false);

			if (onFinish) onFinish();
		});
	};

	useEffect(() => {
		getDatasets().then((d) => {
			setDatasets(d);
		});
	}, []);

	return (
		<div className="w-full h-full">
			{!uploaded && !file && !loading
				? (
					<UploadFile setFile={setFile} />
				)
				: null}
			{!uploaded && file && !loading
				? (
					<ChooseDataset
						datasetId={datasetId}
						datasets={datasets}
						handleFinish={handleFinish}
						setDataset={setDataset}
						setDatasets={setDatasets}
						setIsPublic={setIsPublic}
					/>
				)
				: null}
			{loading
				? (
					<LoadingUpload />
				)
				: null}
			{uploaded && !file && !loading
				? (
					<div
						className={`text-oxfordblue border-oxfordblue w-full h-full rounded-md p-4 flex
											flex-col justify-center items-center border-2`}
					>
						<h1 className="text-2xl font-bold">Upload success!</h1>
						<p className="text-xl">
							File uploaded successfully.
						</p>
						<Button className="my-2" onClick={() => { setUploaded(false); }}>
							Upload another file
							<ArrowRightIcon className="w-4 h-4 ml-2" />
						</Button>
					</div>
				)
				: null}
		</div>
	);
};

export default Upload;
