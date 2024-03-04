import React, { useState } from "react";

import Button from "../../basic/Button";
import Modal from "../../basic/Modal";
import createNewGroup from "../../../helpers/api/webApi/group/createNewGroup";
import ErrorToast from "../../../helpers/toast/errorToast";

const NewGroupModal = ({ isOpen, setIsOpen, onClose }) => {
	const [newGroup, setNewGroup] = useState(null);

	const handleCreate = () => {
		createNewGroup(newGroup).then(() => {
			setNewGroup(null);
			onClose();
			setIsOpen(false);
		}).catch((error) => {
			ErrorToast(error);
		});
	};

	return (
		<Modal isOpen={isOpen}>
			<div className="bg-white p-4 rounded-md w-max">
				<h1 className="text-2xl font-bold">Create a new group</h1>
				<input
					className="bg-transparent border-black border-b-2 outline-none my-2"
					placeholder="Group name"
					autoComplete="off"
					value={newGroup}
					onChange={(e) => { setNewGroup(e.target.value); }}
				/>
				<div className="flex items-center">
					<Button className="rounded-full px-2 my-2 mr-2" onClick={handleCreate}>
						Create
					</Button>
					<Button
						className="rounded-full px-2 my-2 bg-offwhite !text-indianred"
						onClick={() => { setIsOpen(false); }}
					>
						Cancel
					</Button>
				</div>
			</div>
		</Modal>
	);
};

export default NewGroupModal;
