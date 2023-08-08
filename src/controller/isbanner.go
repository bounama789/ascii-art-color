package controller

func IsBanner(term string) bool {

	flags := []string{
		"alligator-italic", "alligator", "alphabet", "arrows", "avatar", "banner-3d", "banner",
		"cards", "graceful", "lil-devil", "ogre", "rectangles", "shadow", "slant",
		"small", "soft", "standard", "stars-wars", "thinkertoy", "train", "varsiry",
	}

	for _, flag := range flags {
		if term == flag {
			return true
		}
	}

	return false
}
