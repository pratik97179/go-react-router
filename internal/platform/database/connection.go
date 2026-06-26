package database

import "github.com/supabase-community/supabase-go"

func New(URL string, Key string) (*supabase.Client, error) {
	db, err := supabase.NewClient(URL, Key, nil)

	if err != nil {
		return nil, err
	}

	return db, nil
}
