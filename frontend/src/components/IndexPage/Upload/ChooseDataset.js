import React, { useEffect, useState } from "react";
import { ArrowRightIcon } from "@heroicons/react/24/solid";

import Button from "../../basic/Button";
import NewDatasetModal from "./NewDatasetModal";

const ChooseDataset = ({
	datasets, setDataset, setDatasets, handleFinish, datasetId,
}) => {
	const [showNewDatasetModal, setShowNewDatasetModal] = useState(false);

	useEffect(() => {
		if (datasetId) {
			setDataset(datasetId);
		}
	}, []);

	return (
		<div
			className={`text-oxfordblue border-oxfordblue w-full h-full rounded-md p-4 flex
											flex-col justify-center items-center border-2`}
		>
			<NewDatasetModal
				isOpen={showNewDatasetModal}
				setDataset={setDataset}
				setDatasets={setDatasets}
			/>
			<h1 className="text-2xl font-bold">File selected!</h1>
			{datasetId
				? null
				: (
					<>
						<p className="text-xl">
							Choose which dataset to put the file in:
						</p>
						<select
							id="field"
							className="text-black px-3 py-2 rounded-3xl border-2 mx-2 my-2"
							onChange={(e) => {
								console.log(e.target.value);
								if (e.target.value === "new") {
									setShowNewDatasetModal(true);
								}
								setDataset(e.target.value);
							}}
							aria-label="choose dataset"
						>
							<option hidden default>Choose dataset</option>
							{datasets.map((d) => {
								return (<option key={d.id} value={d.id}>{d.dataset_name}</option>);
							})}
							<hr />
							<option value="new"> + New dataset</option>
						</select>
					</>
				)}
			<Button className="my-2" onClick={handleFinish}>
				Upload
				<ArrowRightIcon className="w-4 h-4 ml-2" />
			</Button>
		</div>
	);
};

export default ChooseDataset;
