import React, { useEffect, useState } from "react";
import Button from "../../basic/Button";
import Modal from "../../basic/Modal";
import createDataset from "../../../helpers/api/webApi/dataset/createDataset";

const NewDatasetModal = ({ isOpen, setDataset, setDatasets }) => {
	const [newDataset, setNewDataset] = useState(null);
	const [showModal, setShowModal] = useState(isOpen);

	const handleCreateNewDataset = () => {
		createDataset(newDataset).then((d) => {
			setDataset(d.id);
			setDatasets((prev) => { return [...prev, d]; });
			setNewDataset(null);
			setShowModal(false);
		});
	};
	const handleCancel = () => {
		setNewDataset(null);
		setShowModal(false);
	};

	useEffect(() => {
		setShowModal(isOpen);
	}, [isOpen]);

	return (
		<Modal isOpen={showModal}>
			<div className="bg-white p-4 rounded-md w-max">
				<h1 className="text-2xl font-bold">Create a new dataset</h1>
				<input
					className="bg-transparent border-black border-b-2 outline-none my-2"
					placeholder="Dataset name"
					autoComplete="off"
					value={newDataset}
					onChange={(e) => { setNewDataset(e.target.value); }}
				/>
				<div className="flex items-center">
					<Button className="rounded-full px-2 my-2 mr-2" onClick={handleCreateNewDataset}>
						Create
					</Button>
					<Button
						className="rounded-full px-2 my-2 bg-offwhite !text-indianred"
						onClick={handleCancel}
					>
						Cancel
					</Button>
				</div>
			</div>
		</Modal>
	);
};

export default NewDatasetModal;
