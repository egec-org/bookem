package org.everly.apps.game.finder.service.rest.model;

import java.util.Date;
import java.util.HashMap;
import java.util.Map;
import java.util.Set;
import java.util.TreeSet;
import java.util.stream.Collectors;

public class Event {

	private String name;
	private Date startDate;
	private Date endDate;
	private Map<GuestUser, RSVP> attendees;
	private Set<String> matchingInterests;

	public Event(final String name, final Date startDate, final Date endDate) {
		this.name = name;
		this.startDate = startDate;
		this.endDate = endDate;
		this.attendees = new HashMap<>();
		this.matchingInterests = new TreeSet<>();
	}

	public String getName() {
		return name;
	}

	public void setName(final String name) {
		this.name = name;
	}

	public Date getStartDate() {
		return startDate;
	}

	public void setStartDate(final Date startDate) {
		this.startDate = startDate;
	}

	public Date getEndDate() {
		return endDate;
	}

	public void setEndDate(final Date endDate) {
		this.endDate = endDate;
	}

	public Map<GuestUser, RSVP> getAttendees() {
		return attendees;
	}

	public void mergeAttendeeRSVP(final GuestUser attendee, final RSVP rsvp) {
		if (attendee != null && rsvp != null) {
			this.attendees.put(attendee, rsvp);
		}
	}

	public Set<String> getMatchingInterests() {
		return matchingInterests;
	}

	public void setMatchingInterests(final Set<String> matchingInterests) {
		this.matchingInterests.clear();
		if (matchingInterests == null) {
			return;
		}
		this.matchingInterests.addAll(matchingInterests.parallelStream().map(String::toLowerCase).collect(Collectors.toSet()));
	}

	public void addMatchingInterest(final String matchingInterest) {
		if (matchingInterest != null && !matchingInterest.trim().isEmpty()) {
			this.matchingInterests.add(matchingInterest.toLowerCase());
		}
	}
}
