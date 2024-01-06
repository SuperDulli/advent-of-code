/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   get_next_line_utils.c                              :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: chelmerd <chelmerd@student.wolfsburg.de    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/29 13:42:40 by chelmerd          #+#    #+#             */
/*   Updated: 2021/12/02 11:46:05 by chelmerd         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include "get_next_line.h"

/*
* Description
*	Allocates (with malloc(3)) and returns a new string, which is the result of
*	the concatenation of ’s1’ and ’s2’.
*
* Parameters
*	#1. The prefix string.
*	#2. The suffix string.
*
* Return Values
*	The new string.
*	NULL if the allocation fails.
*/
char	*ft_strjoin(const char *s1, const char *s2)
{
	char	*str;
	size_t	len;
	size_t	i;

	if (!s1 || !s2)
		return (NULL);
	len = ft_strlen(s1) + ft_strlen(s2);
	str = (char *) malloc(sizeof (char) * (len + 1));
	i = 0;
	if (str == NULL)
		return (NULL);
	while (*s1)
	{
		str[i++] = *s1;
		s1++;
	}
	while (*s2)
	{
		str[i++] = *s2;
		s2++;
	}
	str[i] = '\0';
	return (str);
}

/*
* Description
*	Computes the length of the string s.
*	The strnlen() function attempts to compute the length of s, but never scans
*	beyond the first maxlen bytes of s.
*
* Return Values
*	The number of characters that precede the terminating NUL character. 
*/
size_t	ft_strlen(const char *s)
{
	size_t	len;

	len = 0;
	while (*s)
	{
		len++;
		s++;
	}
	return (len);
}

/*
* Description
*	Locates the first occurrence of c (converted to a char) in the string pointed
*	to by s.  The terminating null character is considered to be part of the
*	string; therefore if c is \\0, the functions locate the terminating \\0.
*
* Return Values
*	A pointer to the located character,
*	or NULL if the character does not appear in the string.
*/
char	*ft_strchr(const char *s, int c)
{
	while (*s)
	{
		if (*s == (char) c)
			return ((char *) s);
		s++;
	}
	if (*s == (char) c)
		return ((char *) s);
	else
		return (NULL);
}

/*
* Description
*	Allocates sufficient memory for a copy of the string s1, does the copy,
*	and returns a pointer to it.  The pointer may subsequently be used as an
*	argument to the function free(3).
*
* Return Values
*	A pointer to the duplicate of the string, or
*	NULL if the Allocation fails.
*/
char	*ft_strdup(const char *s)
{
	size_t	len;
	char	*dup;

	len = ft_strlen(s);
	dup = (char *) malloc(sizeof (char) * (len + 1));
	if (dup == NULL)
		return (NULL);
	dup[len] = '\0';
	while (len)
	{
		dup[len - 1] = s[len - 1];
		len--;
	}
	return (dup);
}

/*
* Description
*	Allocates (with malloc(3)) and returns a substring from the string ´s´.
*
* Parameters
*	#1. The string from which to create the substring.
*	#2. The start index of the substring in the string ’s’.
*	#3. The maximum length of the substring.
*
* Return Values
*	The substring.
*	NULL if the allocation fails.
*/
char	*ft_substr(char const *s, unsigned int start, size_t len)
{
	char	*sub;
	int		i;

	if (!s)
		return (NULL);
	sub = (char *) malloc(sizeof (char) * (len + 1));
	if (sub == NULL)
		return (NULL);
	if (start >= ft_strlen(s))
	{
		*sub = '\0';
		return (sub);
	}
	i = 0;
	while (len)
	{
		sub[i] = s[start + i];
		if (sub[i] == '\0')
			break ;
		i++;
		len--;
	}
	sub[i] = '\0';
	return (sub);
}
