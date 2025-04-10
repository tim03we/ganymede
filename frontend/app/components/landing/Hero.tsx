"use client"
import { Box, Button, Container, Flex, Group, Title, Text, useMantineTheme } from '@mantine/core';
import classes from './Hero.module.css';
import Image from 'next/image';
import Link from 'next/link';
import { useMediaQuery } from '@mantine/hooks';


export function LandingHero() {
  const theme = useMantineTheme()
  const isMobile = useMediaQuery(`(max-width: ${theme.breakpoints.sm})`);
  return (
    <div className={classes.root}>
      <Container size="xxl">
        <Group justify="space-between">

          <div>
            <Text className={classes.title}>DuckVOD.net</Text>
            <Title c={theme.colors.gray[3]} mt={5} order={3}>Eine Plattform zum Archivieren von Livestreams und Videos im Teich.</Title>
            <Flex mt={10}>
              <Button
                variant="gradient"
                gradient={{ from: 'blue', to: 'purple' }}
                component={Link}
                href="/channels"
                className={classes.button}
              >
                Kanäle
              </Button>
              <Button
                ml={10}
                variant="default"
                component={Link}
                href="/login"
              >
                Login
              </Button>
            </Flex>
          </div>

          {!isMobile && (
            <Box>
              <Flex justify={"center"} align={"center"}>
                <div className={classes.logoBackground}></div>
                <Image src="/images/duckvod_logo.png" height={100} width={100} alt="DuckVOD Logo" className={classes.logo} />
              </Flex>
            </Box>
          )}
        </Group>
        {/* <div className={classes.inner}>
          <div className={classes.content}>
            <Title className={classes.title}>
              A{' '}
              <Text
                component="span"
                inherit
                variant="gradient"
                gradient={{ from: 'pink', to: 'yellow' }}
              >
                fully featured
              </Text>{' '}
              React components library
            </Title>

            <Text className={classes.description} mt={30}>
              Build fully functional accessible web applications with ease – Mantine includes more
              than 100 customizable components and hooks to cover you in any situation
            </Text>

            <Button
              variant="gradient"
              gradient={{ from: 'pink', to: 'yellow' }}
              size="xl"
              className={classes.control}
              mt={40}
            >
              Get started
            </Button>
          </div>
        </div> */}
      </Container>
    </div>
  );
}