import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { ToastContainer } from "react-toastify";
import { BuildingLibraryIcon } from "@heroicons/react/24/outline";
import { UserIcon } from "@heroicons/react/24/solid";

import SearchBar from "../components/SearchBar/SearchBar";

import ErrorPage from "./ErrorPage";
import getUserProfile from "../helpers/api/webApi/user/getUserProfile";
import Dataset from "../components/basic/Dataset";
import Activity from "../components/basic/Activity";

const ProfilePage = () => {
	const [notFound, setNotFound] = useState(false);
	const [profile, setProfile] = useState({});

	const { userId } = useParams();

	useEffect(() => {
		getUserProfile(userId).then((data) => {
			setProfile(data.profile);
		}).catch(() => {
			setNotFound(true);
		});
	}, []);

	return notFound
		? <ErrorPage />
		: (
			<div className="w-screen h-full bg-offwhite">
				<SearchBar />
				<ToastContainer />
				<div className="flex flex-col justify-center items-center w-full">
					<div className="w-full p-8 max-w-5xl w-full flex items-center">
						<div className="min-w-48">
							{profile.avatar === ""
								? (
									<div
										className={`bg-oxfordblue w-48 aspect-square rounded-full flex justify-center 
													items-center`}
									>
										<UserIcon className="h-24 w-24 text-white m-auto" />
									</div>
								) : (
									<img
										className="rounded-full w-48 h-48"
										src={profile.avatar}
										alt="avatar"
									/>
								)}
						</div>
						<div className="ml-8">
							<h1 className="text-3xl font-bold">{profile.name}</h1>
							<div className="flex items-center mb-3 mt-1">
								<BuildingLibraryIcon className="h-6 w-6 text-gray-400" />
								<p className="text-gray-500">University of Leeds</p>
							</div>
							<p>{profile.bio}</p>
						</div>
					</div>
					<div className="w-full p-8 max-w-5xl w-full flex">
						<div className="w-1/2 p-2">
							<h2 className="font-semibold text-2xl pb-4">Recent Activity</h2>
							{profile !== {} && profile.activity
								? profile.activity.map((activity, index) => {
									return (
										<Activity
											activity={activity}
											last={index + 1 === profile.activity.length}
										/>
									);
								})
								: null }
						</div>
						<div className="w-1/2 p-2">
							<h2 className="font-semibold text-2xl px-2">Public Datasets</h2>
							<div className="flex flex-wrap">
								{profile !== {} && profile.datasets
									? profile.datasets.map((dataset) => {
										return (
											<Dataset key={dataset.id} dataset={dataset} />
										);
									})
									: null }
							</div>
						</div>
					</div>
				</div>
			</div>
		);
};

export default ProfilePage;
