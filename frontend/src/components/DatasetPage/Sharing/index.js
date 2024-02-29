import React, { useEffect, useState } from "react";
import { PlusIcon } from "@heroicons/react/24/outline";

import getUserShares from "../../../helpers/api/webApi/datasetSharing/getUserShares";
import getGroupShares from "../../../helpers/api/webApi/datasetSharing/getGroupShares";
import ErrorToast from "../../../helpers/toast/errorToast";
import User from "./User";
import Group from "./Group";
import UserShareModal from "./UserShareModal";
import removeUserShare from "../../../helpers/api/webApi/datasetSharing/removeUserShare";

const Sharing = ({ datasetId }) => {
	const [users, setUsers] = useState([]);
	const [groups, setGroups] = useState([]);

	const [showUserModal, setShowUserModal] = useState(false);

	const handleRemoveUserShare = (userId) => {
		removeUserShare(datasetId, userId).then(() => {
			getUserShares(datasetId).then((data) => {
				setUsers(data.users);
			}).catch((error) => {
				ErrorToast(error);
			});
		}).catch((error) => {
			ErrorToast(error);
		});
	};

	useEffect(() => {
		getUserShares(datasetId).then((data) => {
			setUsers(data.users);
		}).catch((error) => {
			ErrorToast(error);
		});
		getGroupShares(datasetId).then((data) => {
			setGroups(data.groups);
		}).catch((error) => {
			ErrorToast(error);
		});
	}, []);

	return (
		<div className="w-full">
			<UserShareModal
				isOpen={showUserModal}
				setIsOpen={setShowUserModal}
				datasetId={datasetId}
				onClose={() => {
					getUserShares(datasetId).then((data) => {
						setUsers(data.users);
					}).catch((error) => {
						ErrorToast(error);
					});
				}}
			/>
			<h3 className="text-lg font-semibold">Users shared with</h3>
			<div className="flex flex-wrap w-full">
				{users.map((user) => {
					return (
						<User
							key={user.id}
							name={`${user.user.first_name} ${user.user.last_name}`}
							avatar={user.avatar}
							access={user.write_access ? "write" : "read"}
							onRemove={() => {
								handleRemoveUserShare(user.user.id);
							}}
						/>
					);
				})}
				<button
					type="button"
					aria-label="add user"
					className={`w-44 h-36 border-2 border-oxfordblue-200 border-dashed my-2 mr-4 p-4 rounded-md 
								flex flex-col items-center hover:!underline`}
					onClick={() => {
						setShowUserModal(true);
					}}
				>
					<div className="rounded-full border-2 border-dashed border-gray-400 aspect-square p-4">
						<PlusIcon className="h-6 w-6 text-oxfordblue-400" />
					</div>
					<p className="text-md font-semibold py-2">Share with user</p>
				</button>
			</div>

			<h3 className="text-lg font-semibold mt-8">Groups shared with</h3>
			<div className="flex flex-wrap w-full">
				{groups.map((group) => {
					return (
						<div key={group.id}>
							<p>{group.name}</p>
						</div>
					);
				})}
				<button
					type="button"
					aria-label="add user"
					className={`w-44 h-36 border-2 border-oxfordblue-200 border-dashed my-2 mr-4 p-4 rounded-md 
								flex flex-col items-center hover:!underline`}
				>
					<div className="rounded-full border-2 border-dashed border-gray-400 aspect-square p-4">
						<PlusIcon className="h-6 w-6 text-oxfordblue-400" />
					</div>
					<p className="text-md font-semibold py-2">Share with Group</p>
				</button>
			</div>
		</div>
	);
};

export default Sharing;
