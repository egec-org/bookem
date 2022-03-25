package org.everly.apps.game.finder.service.rest.model;

import java.util.Set;
import java.util.TreeSet;
import java.util.stream.Collectors;

public class GuestUser {

	private final String username;
	private final Set<String> interests;

	public GuestUser(final String username) {
		if (username == null || username.trim().isEmpty()) {
			throw new IllegalArgumentException("Username cannot be null or empty");
		}
		this.username = username.trim();
		this.interests = new TreeSet<>();
	}

	public String getUsername() {
		return username;
	}

	public Set<String> getInterests() {
		return interests;
	}

	public void setInterests(final Set<String> interests) {
		this.interests.clear();
		if (interests == null) {
			return;
		}
		this.interests.addAll(interests.parallelStream().map(String::toLowerCase).collect(Collectors.toSet()));
	}

	public void addInterest(final String interest) {
		this.interests.add(interest.toLowerCase());
	}

	public void removeInterest(final String interest) {
		this.interests.remove(interest.toLowerCase());
	}

	@Override
	public boolean equals(final Object other) {
		if (other == null || !(other instanceof GuestUser)) {
			return false;
		}

		GuestUser otherUser = (GuestUser) other;
		return this.username.equalsIgnoreCase(otherUser.username);
	}
}
