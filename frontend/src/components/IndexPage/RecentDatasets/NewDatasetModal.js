import React, { useState } from "react";
import Button from "../../basic/Button";
import Modal from "../../basic/Modal";
import createDataset from "../../../helpers/api/webApi/dataset/createDataset";
import ErrorToast from "../../../helpers/toast/errorToast";

const NewDatasetModal = ({ isOpen, setIsOpen, onClose }) => {
	const [newDataset, setNewDataset] = useState("");

	const handleCreateNewDataset = () => {
		createDataset(newDataset).then(() => {
			setNewDataset(null);
			setIsOpen(false);
			onClose();
		}).catch((error) => {
			ErrorToast(error);
		});
	};
	const handleCancel = () => {
		setNewDataset("");
		setIsOpen(false);
	};

	return (
		<Modal isOpen={isOpen}>
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
