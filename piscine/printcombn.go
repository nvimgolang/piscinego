package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n == 1 {
		for i := '0'; i <= '9'; i++ {
			z01.PrintRune(i)
			if i != '9' {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	} else if n == 2 {
		for i := '0'; i <= '9'; i++ {
			for j := i + 1; j <= '9'; j++ {
				z01.PrintRune(i)
				z01.PrintRune(j)
				if i != '8' || j != '9' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	} else if n == 3 {
		for i := '0'; i <= '9'; i++ {
			for j := i + 1; j <= '9'; j++ {
				for l := j + 1; l <= '9'; l++ {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(l)
					if i != '7' || j != '8' || l != '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	} else if n == 4 {
		for i := '0'; i <= '9'; i++ {
			for j := i + 1; j <= '9'; j++ {
				for l := j + 1; l <= '9'; l++ {
					for m := l + 1; m <= '9'; m++ {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(l)
						z01.PrintRune(m)
						if i != '6' || j != '7' || l != '8' || m != '9' {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	} else if n == 5 {
		for i := '0'; i <= '9'; i++ {
			for j := i + 1; j <= '9'; j++ {
				for l := j + 1; l <= '9'; l++ {
					for m := l + 1; m <= '9'; m++ {
						for p := m + 1; p <= '9'; p++ {
							z01.PrintRune(i)
							z01.PrintRune(j)
							z01.PrintRune(l)
							z01.PrintRune(m)
							z01.PrintRune(p)
							if i != '5' || j != '6' || l != '7' || m != '8' || p != '9' {
								z01.PrintRune(',')
								z01.PrintRune(' ')
							}
						}
					}
				}
			}
		}
	} else if n == 6 {
		for i := '0'; i <= '9'; i++ {
			for j := i + 1; j <= '9'; j++ {
				for l := j + 1; l <= '9'; l++ {
					for m := l + 1; m <= '9'; m++ {
						for p := m + 1; p <= '9'; p++ {
							for q := p + 1; q <= '9'; q++ {
								z01.PrintRune(i)
								z01.PrintRune(j)
								z01.PrintRune(l)
								z01.PrintRune(m)
								z01.PrintRune(p)
								z01.PrintRune(q)
								if i != '4' || j != '5' || l != '6' || m != '7' || p != '8' || q != '9' {
									z01.PrintRune(',')
									z01.PrintRune(' ')
								}
							}
						}
					}
				}
			}
		}
	} else if n == 7 {
		for i := '0'; i <= '9'; i++ {
			for j := i + 1; j <= '9'; j++ {
				for l := j + 1; l <= '9'; l++ {
					for m := l + 1; m <= '9'; m++ {
						for p := m + 1; p <= '9'; p++ {
							for q := p + 1; q <= '9'; q++ {
								for r := q + 1; r <= '9'; r++ {
									z01.PrintRune(i)
									z01.PrintRune(j)
									z01.PrintRune(l)
									z01.PrintRune(m)
									z01.PrintRune(p)
									z01.PrintRune(q)
									z01.PrintRune(r)
									if i != '3' || j != '4' || l != '5' || m != '6' || p != '7' || q != '8' || r != '9' {
										z01.PrintRune(',')
										z01.PrintRune(' ')
									}
								}
							}
						}
					}
				}
			}
		}
	} else if n == 8 {
		for i := '0'; i <= '9'; i++ {
			for j := i + 1; j <= '9'; j++ {
				for l := j + 1; l <= '9'; l++ {
					for m := l + 1; m <= '9'; m++ {
						for p := m + 1; p <= '9'; p++ {
							for q := p + 1; q <= '9'; q++ {
								for r := q + 1; r <= '9'; r++ {
									for s := r + 1; s <= '9'; s++ {
										z01.PrintRune(i)
										z01.PrintRune(j)
										z01.PrintRune(l)
										z01.PrintRune(m)
										z01.PrintRune(p)
										z01.PrintRune(q)
										z01.PrintRune(r)

										z01.PrintRune(s)
										if i != '2' || j != '3' || l != '4' || m != '5' || p != '6' || q != '7' || r != '8' || s != '9' {
											z01.PrintRune(',')
											z01.PrintRune(' ')
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else if n == 9 {
		for i := '0'; i <= '9'; i++ {
			for j := i + 1; j <= '9'; j++ {
				for l := j + 1; l <= '9'; l++ {
					for m := l + 1; m <= '9'; m++ {
						for p := m + 1; p <= '9'; p++ {
							for q := p + 1; q <= '9'; q++ {
								for r := q + 1; r <= '9'; r++ {
									for s := r + 1; s <= '9'; s++ {
										for t := s + 1; t <= '9'; t++ {
											z01.PrintRune(i)
											z01.PrintRune(j)
											z01.PrintRune(l)
											z01.PrintRune(m)
											z01.PrintRune(p)
											z01.PrintRune(q)
											z01.PrintRune(r)
											z01.PrintRune(s)
											z01.PrintRune(t)
											if i != '1' || j != '2' || l != '3' || m != '4' || p != '5' || q != '6' || r != '7' || s != '8' || t != '9' {
												z01.PrintRune(',')
												z01.PrintRune(' ')
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else if n == 10 {
		for i := '0'; i <= '9'; i++ {
			for j := i + 1; j <= '9'; j++ {
				for l := j + 1; l <= '9'; l++ {
					for m := l + 1; m <= '9'; m++ {
						for p := m + 1; p <= '9'; p++ {
							for q := p + 1; q <= '9'; q++ {
								for r := q + 1; r <= '9'; r++ {
									for s := r + 1; s <= '9'; s++ {
										for t := s + 1; t <= '9'; t++ {
											for u := t + 1; u <= '9'; u++ {
												z01.PrintRune(i)
												z01.PrintRune(j)
												z01.PrintRune(l)
												z01.PrintRune(m)
												z01.PrintRune(p)
												z01.PrintRune(q)
												z01.PrintRune(r)
												z01.PrintRune(s)
												z01.PrintRune(t)
												z01.PrintRune(u)
												if i != '0' || j != '1' || l != '2' || m != '3' || p != '4' || q != '5' || r != '6' || s != '7' || t != '8' || u != '9' {
													z01.PrintRune(',')
													z01.PrintRune(' ')
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
