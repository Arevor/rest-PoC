create table cards
(
	cardid serial not null
		constraint cards_pkey
			primary key,
	title text not null,
	cardtype text not null,
	desc1 text,
	desc2 text,
	desc3 text,
	priority integer default 0 not null,
	severity integer default 0 not null,
	assignedto text,
	casenumber integer,
	boardcol integer default 0 not null
)
;

create unique index cards_cardid_uindex
	on cards (cardid)
;

create table comments
(
	commentid serial not null
		constraint comments_pkey
			primary key,
	cardid integer not null
		constraint comments_cards_cardid_fk
			references cards,
	author text not null,
	message text not null,
	posted timestamp default CURRENT_TIMESTAMP
)
;

create unique index comments_commentid_uindex
	on comments (commentid)
;

