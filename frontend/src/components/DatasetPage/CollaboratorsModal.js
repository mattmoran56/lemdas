import React, { useEffect, useState } from "react";
import { TrashIcon } from "@heroicons/react/24/outline";

import Modal from "../basic/Modal";
import ErrorToast from "../../helpers/toast/errorToast";
import getCollaborators from "../../helpers/api/webApi/datasetCollaborators/getCollaborators";
import SearchUser from "../basic/SearchUser";
import Button from "../basic/Button";
import deleteCollaborator from "../../helpers/api/webApi/datasetCollaborators/deleteCollaborator";
import addCollaborator from "../../helpers/api/webApi/datasetCollaborators/addCollaborator";

const CollaboratorsModal = ({
	isOpen, setIsOpen, onClose, datasetId,
}) => {
	const [collaborators, setCollaborators] = useState([]);

	const handleAdd = (user) => {
		addCollaborator(datasetId, user.id).then(() => {
			getCollaborators(datasetId).then((data) => {
				setCollaborators(data.collaborators);
			}).catch((error) => {
				setIsOpen(false);
				ErrorToast(error);
			});
		}).catch((error) => {
			ErrorToast(error);
		});
	};

	const handleDelete = (userId) => {
		deleteCollaborator(datasetId, userId).then(() => {
			getCollaborators(datasetId).then((data) => {
				setCollaborators(data.collaborators);
			}).catch((error) => {
				setIsOpen(false);
				ErrorToast(error);
			});
		}).catch((error) => {
			ErrorToast(error);
		});
	};

	useEffect(() => {
		getCollaborators(datasetId).then((data) => {
			setCollaborators(data.collaborators);
		}).catch((error) => {
			setIsOpen(false);
			ErrorToast(error);
		});
	}, []);

	return (
		<Modal isOpen={isOpen}>
			<div className="bg-white p-4 rounded-md w-max">
				<h1 className="text-2xl font-bold">Collaborators</h1>
				<div className="flex items-center mt-8">
					<SearchUser setUser={handleAdd} />
				</div>
				<div className="w-full mt-8">
					<h2 className="text-lg font-semibold mb-2">Current Collaborators</h2>
					{collaborators.map((collaborator, i) => {
						return (
							<div
								key={collaborator.id}
								className={`flex justify-between items-center p-3 ${i % 2 === 0 ? "bg-gray-100" : ""}`}
							>
								<p>
									{collaborator.user.first_name} {collaborator.user.last_name}
									({collaborator.user.email})
								</p>
								<button
									type="button"
									aria-label="remove collaborator"
									className="text-red-500"
									onClick={() => {
										handleDelete(collaborator.user.id);
									}}
								>
									<TrashIcon className="h-6 w-6 mr-3 hover:text-red-500" />
								</button>
							</div>
						);
					})}
				</div>
			</div>
			<Button
				className="mt-4 ml-4"
				onClick={() => {
					onClose();
					setIsOpen(false);
				}}
			>
				Done
			</Button>
		</Modal>
	);
};

export default CollaboratorsModal;
