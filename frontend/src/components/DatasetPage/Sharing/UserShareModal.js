import React from "react";
import Modal from "../../basic/Modal";
import SearchUser from "../../basic/SearchUser";
import Button from "../../basic/Button";
import shareWithUser from "../../../helpers/api/webApi/datasetSharing/shareWithUser";
import ErrorToast from "../../../helpers/toast/errorToast";

const UserShareModal = ({
	isOpen, setIsOpen, datasetId, onClose,
}) => {
	const [user, setUser] = React.useState(null);
	const [access, setAccess] = React.useState("read");
	const handleSelectUser = (selectedUser) => {
		setUser(selectedUser);
	};

	const handleShare = () => {
		shareWithUser(datasetId, user.id, access === "write").then(() => {
			onClose();
			setIsOpen(false);
		}).catch((error) => {
			ErrorToast(error);
		});
	};

	return (
		<Modal isOpen={isOpen}>
			<div className="bg-white p-4 rounded-md w-max">
				<h1 className="text-2xl font-bold">Share with User</h1>
				<div className="flex items-center mt-8">
					<SearchUser setUser={handleSelectUser} />
				</div>
				<select
					value={access}
					className="border-2 rounded-md mb-0 mt-4"
					onChange={(e) => { setAccess(e.target.value); }}
				>
					<option value="read">Read</option>
					<option value="write">Write</option>
				</select>
			</div>
			<div className="w-full flex">
				<Button
					className="mt-4 ml-4"
					onClick={() => {
						handleShare();
					}}
				>
					Share
				</Button>
				<Button
					className="mt-4 ml-4 bg-transparent !text-indianred"
					onClick={() => {
						onClose();
						setIsOpen(false);
					}}
				>
					Close
				</Button>
			</div>
		</Modal>
	);
};

export default UserShareModal;
