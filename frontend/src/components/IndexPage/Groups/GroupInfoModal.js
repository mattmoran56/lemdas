import React, { useEffect, useState } from "react";
import { TrashIcon } from "@heroicons/react/24/outline";

import Modal from "../../basic/Modal";
import ErrorToast from "../../../helpers/toast/errorToast";
import SearchUser from "../../basic/SearchUser";
import Button from "../../basic/Button";
import getGroupMembers from "../../../helpers/api/webApi/group/getGroupMembers";
import addUserToGroup from "../../../helpers/api/webApi/group/addUserToGroup";
import deleteGroupMember from "../../../helpers/api/webApi/group/removeGroupMember";
import useAuth from "../../../hooks/useAuth";
import deleteGroup from "../../../helpers/api/webApi/group/deleteGroup";

const GroupInfoModal = ({
	isOpen, setIsOpen, group, onClose,
}) => {
	const [members, setMembers] = useState([]);
	const [user, setUser] = useState({});

	const { id } = useAuth();

	const handleAdd = () => {
		addUserToGroup(group.id, user.id).then(() => {
			getGroupMembers(group.id).then((data) => {
				setMembers(data.members);
				setUser({});
			}).catch((error) => {
				ErrorToast(error);
			});
		}).catch((error) => {
			ErrorToast(error);
		});
	};

	const handleDelete = (userId) => {
		deleteGroupMember(group.id, userId).then(() => {
			getGroupMembers(group.id).then((data) => {
				setMembers(data.members);
			}).catch((error) => {
				ErrorToast(error);
			});
		}).catch((error) => {
			ErrorToast(error);
		});
	};

	useEffect(() => {
		getGroupMembers(group.id).then((data) => {
			setMembers(data.members);
		}).catch((error) => {
			onClose();
			setIsOpen(false);
			ErrorToast(error);
		});
	}, []);

	return (
		<Modal isOpen={isOpen}>
			<div className="bg-white p-4 rounded-md w-max">
				<h1 className="text-2xl font-bold">{group.group_name}</h1>
				<div className="flex items-center mt-8">
					<SearchUser key={user.id} setUser={setUser} />
					<Button
						className="ml-4"
						onClick={handleAdd}
						disabled={group.owner_id !== id}
					>
						Add
					</Button>
				</div>
				<div className="w-full mt-8">
					<h2 className="text-lg font-semibold mb-2">Current Members</h2>
					{members.map((member, i) => {
						return (
							<div
								key={member.id}
								className={`flex justify-between items-center p-3 ${i % 2 === 0 ? "bg-gray-100" : ""}`}
							>
								<p>
									{member.first_name} {member.last_name}
									({member.email})
								</p>
								{group.owner_id === id
									? (
										<button
											type="button"
											aria-label="remove collaborator"
											className="text-red-500"
											onClick={() => {
												handleDelete(member.id);
											}}
										>
											<TrashIcon className="h-6 w-6 mr-3 hover:text-red-500" />
										</button>
									) : null}
							</div>
						);
					})}
				</div>
			</div>
			<div className="flex">
				<Button
					className="mt-4 ml-4"
					onClick={() => {
						onClose();
						setIsOpen(false);
					}}
				>
					Done
				</Button>
				<Button
					onClick={() => {
						deleteGroup(group.id).then(() => {
							onClose();
							setIsOpen(false);
						}).catch((error) => {
							ErrorToast(error);
						});
					}}
					className="!text-indianred bg-transparent mt-4 ml-4"
				>
					<TrashIcon className="h-6 w-6 mr-3" />
					Delete Group
				</Button>
			</div>
		</Modal>
	);
};

export default GroupInfoModal;
