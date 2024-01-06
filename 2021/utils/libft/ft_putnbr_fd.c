/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_putnbr_fd.c                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: chelmerd <chelmerd@student.wolfsburg.de    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/17 12:47:55 by chelmerd          #+#    #+#             */
/*   Updated: 2021/11/26 12:33:37 by chelmerd         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include <unistd.h>

void		ft_putnbr_fd(int n, int fd);

static void	putnbr(int n, int fd);
static int	ft_abs(int n);

/*
* Description
*	Outputs the integer ’n’ to the given file descriptor.
*
* Parameters
*	#1. The integer to output.
*	#2. The file descriptor on which to write.
*/
void	ft_putnbr_fd(int n, int fd)
{
	if (n == 0)
		write(fd, "0", 1);
	else
	{
		if (n < 0)
			write(fd, "-", 1);
		putnbr(n, fd);
	}
}

/*
* Description
*	Recursivly calls itself to write digit by digit to the given file descriptor.
*/
static void	putnbr(int n, int fd)
{
	char	c;

	if (n != 0)
	{
		if (n > 9 || n < -9)
			putnbr(n / 10, fd);
		c = ft_abs(n % 10) + '0';
		write(fd, &c, 1);
	}
}

static int	ft_abs(int n)
{
	if (n < 0)
		return (-n);
	else
		return (n);
}
